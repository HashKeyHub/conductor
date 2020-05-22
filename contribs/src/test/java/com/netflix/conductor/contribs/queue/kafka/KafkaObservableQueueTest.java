package com.netflix.conductor.contribs.queue.kafka;

import com.netflix.conductor.core.events.queue.Message;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.junit.Test;
import reactor.core.publisher.Flux;
import reactor.kafka.receiver.ReceiverRecord;

import java.util.LinkedList;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

import static org.junit.Assert.assertEquals;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

public class KafkaObservableQueueTest {

	@Test
	public void testObserve() {
		String value = "{ \"some_key\": \"valuex\", \"externalId\": \"{\\\"taskRefName\\\":\\\"wait_msg\\\",\\\"workflowId\\\":\\\"b0a75e2c-11d2-46ea-82fc-6e87a2905343\\\"}\" }";
		KafkaObservableQueue queue = mock(KafkaObservableQueue.class);
		List<ReceiverRecord<String, String>> records = IntStream.range(0, 10)
				.boxed()
				.map(i -> new ReceiverRecord<>(new ConsumerRecord<>("conductor_kafka_notify_COMPLETED", 1, i, i.toString(), value), null))
				.collect(Collectors.toList());
		when(queue.receive()).thenReturn(Flux.fromIterable(records));
		when(queue.observe()).thenCallRealMethod();

		List<Message> found = new LinkedList<>();
		queue.observe().subscribe(found::add);
		assertEquals(records.size(), found.size());
		for (int i = 0 ; i< records.size(); i++) {
			assertEquals(records.get(i).value(), found.get(i).getPayload());
		}
	}

}
