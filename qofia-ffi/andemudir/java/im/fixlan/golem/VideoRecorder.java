package im.fixlan.golem;

import android.media.MediaRecorder;
import android.view.*;


import java.io.IOException;
import java.io.FileDescriptor;
import java.lang.reflect.Field;

public class VideoRecorder
{
    public static MediaRecorder recorder = null;

    public static void startRec(int fd) {
        recorder = new MediaRecorder();

        recorder.setAudioSource(MediaRecorder.AudioSource.MIC);
        recorder.setVideoSource(MediaRecorder.VideoSource.CAMERA);

        recorder.setOutputFormat(MediaRecorder.OutputFormat.THREE_GPP);

        recorder.setAudioEncoder(MediaRecorder.AudioEncoder.AMR_NB);
        recorder.setVideoEncoder(MediaRecorder.VideoEncoder.H264);

        recorder.setVideoSize(720, 480);
        recorder.setVideoFrameRate(15); //  Requested frame rate (20) is not supported: 15,24,30
        recorder.setOutputFile("/sdcard/jmr.dat");
        // recorder.setOutputFile(fd);
        try {
            recorder.prepare();
        } catch (IOException ex){
        }
        recorder.start();   // Recording is now started
    }

    public static void stopRec(){
        // ...
        recorder.stop();
        recorder.reset();   // You can reuse the object by going back to setAudioSource() step
        recorder.release(); // Now the object cannot be reused
        recorder = null;
    }

    public static void setfd(FileDescriptor fd, int fdi){
        // descriptor
    }

    public static long fileno(FileDescriptor fd) throws IOException {
        try {
            if (fd.valid()) {
                // windows builds use long handle
                long fileno = getFileDescriptorField(fd, "handle", false);
                if (fileno != -1) {
                    return fileno;
                }
                // unix builds use int fd
                return getFileDescriptorField(fd, "fd", true);
            }
        } catch (IllegalAccessException e) {
            throw new IOException("unable to access handle/fd fields in FileDescriptor", e);
        } catch (NoSuchFieldException e) {
            throw new IOException("FileDescriptor in this JVM lacks handle/fd fields", e);
        }
        return -1;
    }

    private static long getFileDescriptorField(FileDescriptor fd, String fieldName, boolean isInt) throws NoSuchFieldException, IllegalAccessException {
        Field field = FileDescriptor.class.getDeclaredField(fieldName);
        field.setAccessible(true);
        long value = isInt ? field.getInt(fd) : field.getLong(fd);
        field.setAccessible(false);
        return value;
    }

    private static void setFileDescriptorField(FileDescriptor fd, String fieldName, int fdi, boolean isInt) throws NoSuchFieldException, IllegalAccessException {
        Field field = FileDescriptor.class.getDeclaredField(fieldName);
        field.setAccessible(true);
        if (isInt) {
            field.setInt(fd, fdi);
        }else{
            // field.setLong(long(fdi));
        }

        field.setAccessible(false);
        // return value;
    }

}

