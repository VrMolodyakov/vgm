package com.vgm.image.thumbnailService;

import net.coobird.thumbnailator.Thumbnails;
import net.coobird.thumbnailator.name.Rename;

import java.io.File;
import java.io.IOException;

public class ThumbnailServiceImpl implements ThumbnailService {
    @Override
    public void Resize(String path,int w, int h) {
        try {
            Thumbnails.of(new File(path).listFiles())
                    .size(160, 160)
                    .outputFormat("jpg")
                    .toFiles(Rename.PREFIX_DOT_THUMBNAIL);
        }catch (IOException e) {
            e.printStackTrace();
        }
    }
}
