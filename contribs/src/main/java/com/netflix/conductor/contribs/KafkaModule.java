package com.netflix.conductor.contribs;

import com.google.inject.AbstractModule;
import com.google.inject.Provides;
import com.netflix.conductor.common.metadata.tasks.Task;
import com.netflix.conductor.contribs.queue.QueueManager;
import com.netflix.conductor.contribs.queue.kafka.KafkaObservableQueue;
import com.netflix.conductor.core.config.Configuration;
import com.netflix.conductor.core.events.queue.ObservableQueue;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.HashMap;
import java.util.Map;


/**
 * configuration file:
 *
 * <pre>
 * conductor.additional.modules = com.netflix.conductor.contribs.KafkaModule
 * </pre>
 *
 */
public class KafkaModule extends AbstractModule {
	private static final Logger logger = LoggerFactory.getLogger(KafkaModule.class);

	@Override
	protected void configure() {
		bind(QueueManager.class).asEagerSingleton();
		logger.info("Kafka module configured...");
	}

	@Provides
	public Map<Task.Status, ObservableQueue> getQueues(Configuration config) {
		Task.Status[] statuses = new Task.Status[]{Task.Status.COMPLETED, Task.Status.FAILED};
		Map<Task.Status, ObservableQueue> queues = new HashMap<>();
		for (Task.Status status : statuses) {
			String queueName = config.getProperty("workflow.listener.queue.prefix", config.getAppId() + "_kafka_notify_" + status.name());
			String bootstrapServers = config.getProperty("workflow.event.queues.kafka.bootstrap_servers", "127.0.0.1:9092");
			queues.put(status, new KafkaObservableQueue(queueName, bootstrapServers));
		}
		return queues;
	}

}
