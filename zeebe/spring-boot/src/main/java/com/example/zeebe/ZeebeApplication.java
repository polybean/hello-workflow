package com.example.zeebe;

import io.zeebe.client.api.response.ActivatedJob;
import io.zeebe.client.api.worker.JobClient;
import io.zeebe.spring.client.EnableZeebeClient;
import io.zeebe.spring.client.annotation.ZeebeWorker;
import lombok.extern.slf4j.Slf4j;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import java.time.Instant;

@SpringBootApplication
@EnableZeebeClient
public class ZeebeApplication {

	public static void main(String[] args) {
		SpringApplication.run(ZeebeApplication.class, args);
	}

	@ZeebeWorker(type = "ocr-on-claim-attachments")
	public void ocrOnClaimAttachments(final JobClient client, final ActivatedJob job) {
		logJob(job);
		client.newCompleteCommand(job.getKey()).send().join();
	}

	@ZeebeWorker(type = "create-proposals")
	public void createProposals(final JobClient client, final ActivatedJob job) {
		logJob(job);
		client.newCompleteCommand(job.getKey()).send().join();
	}

	private static void logJob(final ActivatedJob job) {
		System.out.println(job);
	}
}
