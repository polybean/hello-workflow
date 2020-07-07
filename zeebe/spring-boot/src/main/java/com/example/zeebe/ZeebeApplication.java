package com.example.zeebe;

import io.zeebe.client.api.response.ActivatedJob;
import io.zeebe.client.api.worker.JobClient;
import io.zeebe.spring.client.EnableZeebeClient;
import io.zeebe.spring.client.annotation.ZeebeWorker;

import java.time.Instant;

import lombok.extern.slf4j.Slf4j;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@EnableZeebeClient
public class ZeebeApplication {

    public static void main(String[] args) {
        SpringApplication.run(ZeebeApplication.class, args);
    }

    @ZeebeWorker(type = "collect-money")
    public void collectMoney(final JobClient client, final ActivatedJob job) {
        logJob(job);
        client.newCompleteCommand(job.getKey()).send().join();
    }

    @ZeebeWorker(type = "fetch-items")
    public void fetchItems(final JobClient client, final ActivatedJob job) {
        logJob(job);
        client.newCompleteCommand(job.getKey()).send().join();
    }

    @ZeebeWorker(type = "ship-parcel")
    public void shipParcel(final JobClient client, final ActivatedJob job) {
        logJob(job);
        client.newCompleteCommand(job.getKey()).send().join();
    }

    private static void logJob(final ActivatedJob job) {
        System.out.println(job);
    }
}
