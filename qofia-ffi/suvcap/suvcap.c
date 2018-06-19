#include <stdlib.h>
#include <unistd.h>

#include <libavcodec/avcodec.h>
#include "libavformat/avformat.h"
#include <libavdevice/avdevice.h>
#include "libavutil/opt.h"

// args: 1: ctx pointer address, integer,
// args: 2: err pointer address, integer
// args: 3: err pointer address, integer
// args: 4: fname pointer address
int main(int argc, char **argv) {

    av_register_all(); // depcreated warning
    avdevice_register_all();
    avcodec_register_all(); // depcreated warning
    avformat_network_init();
    // avfilter_register_all();
    av_log_set_level(AV_LOG_DEBUG);

    long p2 = strtol(argv[2], NULL, 10);
    int *goretp = (int*)(void*)(p2);
    AVFormatContext *avctx = avformat_alloc_context();
    AVDictionary *options = NULL;
    char *fname = argv[4];
    int averr = avformat_open_input(&avctx, fname, NULL, &options);
    if (averr < 0) {
        fprintf(stderr, "open input error1: %d\n", averr);
        // *goretp = averr;
        exit(averr);
    }
    averr = avformat_find_stream_info(avctx, NULL);
    if (averr < 0) {
        fprintf(stderr, "open input error2: %d\n", averr);
        // *goretp = averr;
        exit(averr);
    }

    // success
    // *goretp = 1;
    long p1 = strtol(argv[1], NULL, 10);
    AVFormatContext **goctxp = (AVFormatContext**)(void*)(p1);
    // *goctxp = avctx;
    for (;;) {
        sleep(1);
        continue;
        long p3 = strtol(argv[3], NULL, 10);
        int *gostopp = (int*)(void*)(p3);
        if (*gostopp == 1) {
            avio_close(avctx->pb);
            avformat_free_context(avctx);
            break;
        }
    }
    return 0;
}

