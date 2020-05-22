package com.netflix.conductor.contribs.queue.kafka;

import com.google.common.annotations.VisibleForTesting;
import com.netflix.conductor.core.events.queue.Message;
import com.netflix.conductor.core.events.queue.ObservableQueue;
import org.apache.kafka.clients.consumer.ConsumerConfig;
import org.apache.kafka.common.serialization.StringDeserializer;
import org.reactivestreams.Publisher;
import org.reactivestreams.Subscriber;
import org.reactivestreams.Subscription;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.Exceptions;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Operators;
import reactor.kafka.receiver.KafkaReceiver;
import reactor.kafka.receiver.ReceiverOffset;
import reactor.kafka.receiver.ReceiverOptions;
import reactor.kafka.receiver.ReceiverRecord;
import rx.Observable;
import rx.Producer;

import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.UUID;
import java.util.concurrent.atomic.AtomicIntegerFieldUpdater;

public class KafkaObservableQueue implements ObservableQueue {

	private static Logger logger = LoggerFactory.getLogger(KafkaObservableQueue.class);

	private static final String QUEUE_TYPE = "kafka";

	private String queueName;
	private final ReceiverOptions<String, String> receiverOptions;

	public KafkaObservableQueue(String queueName, String bootstrapServers) {
		this.queueName = queueName;
		Map<String, Object> props = new HashMap<>();
		props.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, bootstrapServers);
		props.put(ConsumerConfig.CLIENT_ID_CONFIG, UUID.randomUUID().toString());
		props.put(ConsumerConfig.GROUP_ID_CONFIG, "sample-group");
		props.put(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class);
		props.put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class);
		props.put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest");
		receiverOptions = ReceiverOptions.create(props);
	}

	@Override
	public Observable<Message> observe() {
		logger.debug("Observe messages");
		Flux<ReceiverRecord<String, String>> receiver = receive();
		return Observable.create(new PublisherAsObservable<>(receiver.map(r -> {
			String id = r.topic() + ":" + r.partition() + ":" + r.offset();
			return new KafkaMessage(id, r.value(), r.receiverOffset());
		})));
	}

	@Override
	public String getType() {
		return QUEUE_TYPE;
	}

	@Override
	public String getName() {
		return queueName;
	}

	@Override
	public String getURI() {
		return queueName;
	}

	@Override
	public List<String> ack(List<Message> messages) {
		messages.stream()
				.filter(m -> m instanceof KafkaMessage)
				.forEach(m -> ((KafkaMessage) m).receiverOffset.acknowledge());
		return Collections.emptyList();
	}

	@Override
	public void publish(List<Message> messages) {
		// not implemented
		throw new UnsupportedOperationException();
	}

	@Override
	public void setUnackTimeout(Message message, long unackTimeout) {
	}

	@Override
	public long size() {
		return 0;
	}

	@VisibleForTesting
	Flux<ReceiverRecord<String, String>> receive() {
		ReceiverOptions<String, String> options = receiverOptions.subscription(Collections.singletonList(queueName))
				.addAssignListener(partitions -> logger.debug("onPartitionsAssigned {}", partitions))
				.addRevokeListener(partitions -> logger.debug("onPartitionsRevoked {}", partitions));
		return KafkaReceiver.create(options).receive();
	}

	private static class PublisherAsObservable<T> implements Observable.OnSubscribe<T> {

		private final Publisher<T> pub;

		private PublisherAsObservable(Publisher<T> pub) {
			this.pub = pub;
		}

		@Override
		public void call(final rx.Subscriber<? super T> subscriber) {
			try {
				pub.subscribe(new SubscriberToRx<>(subscriber));
			} catch (Throwable t) {
				Exceptions.throwIfFatal(t);
				subscriber.onError(t);
			}
		}
	}

	private static class SubscriberToRx<T> implements Subscriber<T>, Producer, Subscription, rx.Subscription {

		private final rx.Subscriber<? super T> subscriber;
		private volatile int terminated;
		private volatile Subscription subscription;

		private static final AtomicIntegerFieldUpdater<SubscriberToRx> TERMINATED =
				AtomicIntegerFieldUpdater.newUpdater(SubscriberToRx.class, "terminated");

		public SubscriberToRx(rx.Subscriber<? super T> subscriber) {
			this.subscriber = subscriber;
			terminated = 0;
		}

		@Override
		public void request(long n) {
			if (n == 0 || isUnsubscribed()) {
				return; // ignore in RxJava
			}
			if (n <= 0L) {
				subscriber.onError(Exceptions.nullOrNegativeRequestException(n));
				return;
			}

			Subscription subscription = this.subscription;
			if (subscription != null) {
				subscription.request(n);
			}
		}

		@Override
		public boolean isUnsubscribed() {
			return terminated == 1;
		}

		@Override
		public void unsubscribe() {
			if (TERMINATED.compareAndSet(this, 0, 1)) {
				Subscription subscription = this.subscription;
				if (subscription != null) {
					this.subscription = null;
					subscription.cancel();
				}
			}
		}

		@Override
		public void cancel() {
			unsubscribe();
		}

		@Override
		public void onSubscribe(final Subscription s) {
			if (Operators.validate(subscription, s)) {
				this.subscription = s;
				subscriber.add(this);
				subscriber.onStart();
				subscriber.setProducer(this);
			}
		}

		@Override
		public void onNext(T o) {
			subscriber.onNext(o);
		}

		@Override
		public void onError(Throwable t) {
			if (TERMINATED.compareAndSet(this, 0, 1)) {
				subscription = null;
				subscriber.onError(t);
			}
		}

		@Override
		public void onComplete() {
			if (TERMINATED.compareAndSet(this, 0, 1)) {
				subscription = null;
				subscriber.onCompleted();
			}
		}
	}

	public static class KafkaMessage extends Message {
		private final ReceiverOffset receiverOffset;

		public KafkaMessage(String id, String payload, ReceiverOffset receiverOffset) {
			super(id, payload, null);
			this.receiverOffset = receiverOffset;
		}
	}
}
