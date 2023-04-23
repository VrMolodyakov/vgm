package com.vgm.image.controller;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.codec.multipart.FilePart;
import org.springframework.web.bind.annotation.RequestPart;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.beans.factory.annotation.Value;
import reactor.core.publisher.Mono;

@RestController
public class ImageController {

    private static Logger LOGGER = LoggerFactory.getLogger(ImageController.class);
    @Value("${image.path}")
    private String uploadDir;

    public Mono<Void> upload(@RequestPart("file-name") String name,
                             @RequestPart("file") Mono<FilePart> fileMonoPar){
        return fileMonoPar.then();
    }

}
