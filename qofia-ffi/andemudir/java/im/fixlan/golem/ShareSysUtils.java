package im.fixlan.golem;

import android.os.*;
import android.content.*;
import android.app.*;

import java.lang.String;
import android.content.Intent;
import java.io.File;
import android.net.Uri;
import android.util.Log;

import java.lang.*;
import java.util.*;
// import java.awt.*; // Cursor
import android.database.Cursor;

/**
 * CdLibsTest     com.march.libs.utils
 * Created by 陈栋 on 16/3/25.
 * 功能:接入系统分享时使用,可以作为接受的一方,也可以作为分享的一方
 */
public class ShareSysUtils {
    public static ShareSysUtils sysShareUtils;
    private Activity activity;
    private OnShareDataOkListener listener;
    public static int TYPE_TEXT = 0, TYPE_IMAGE = 1;

    public ShareSysUtils(Activity activity) {
        this.activity = activity;
    }


    public static ShareSysUtils get(Activity activity) {
        if (sysShareUtils == null) {
            synchronized (ShareSysUtils.class) {
                if (sysShareUtils == null) {
                    sysShareUtils = new ShareSysUtils(activity);
                }
            }
        }
        return sysShareUtils;
    }

    /**
     * 作为接受分享的一方,处理分享来的数据
     *
     * @param listener 处理监听,数据处理好之后会返回
     */
    public void handleShare(OnShareDataOkListener listener) {
        this.listener = listener;
        Intent intent = activity.getIntent();
        String action = intent.getAction();
        String type = intent.getType();
        if (null == action || type == null) {
            Log.e("golemscorner", "没有检索到分享数据");
            return;
        }
        if (Intent.ACTION_SEND.equals(action) && type != null) {
            if ("text/plain".equals(type)) {
                handleSendText(intent);
            } else if (type.startsWith("image/")) {
                handleSendImage(intent);
            }
        } else if (Intent.ACTION_SEND_MULTIPLE.equals(action) && type != null) {
            if (type.startsWith("image/")) {
                handleSendMultipleImages(intent);
            }
        }
    }


    /**
     * 分享文字
     *
     * @param title   文字标题
     * @param content 文字内容
     */
    public void shareText(String title, String content) {
        Intent shareIntent = new Intent();
        shareIntent.setAction(Intent.ACTION_SEND);
        shareIntent.putExtra(Intent.EXTRA_TEXT, content);
        shareIntent.putExtra(Intent.EXTRA_TITLE, title);
        shareIntent.setType("text/plain");
        //设置分享列表的标题，并且每次都显示分享列表
        activity.startActivity(Intent.createChooser(shareIntent, "分享到"));
    }

    /**
     * 分享单张图片
     *
     * @param path 图片的路径
     */
    public void shareSingleImage(String path) {
        //由文件得到uri
        Uri imageUri = Uri.fromFile(new File(path));
//        Log.d("share", "uri:" + imageUri);  //输出：file:///storage/emulated/0/test.jpg
        Intent shareIntent = new Intent();
        shareIntent.setAction(Intent.ACTION_SEND);
        shareIntent.putExtra(Intent.EXTRA_STREAM, imageUri);
        shareIntent.setType("image/*");
        activity.startActivity(Intent.createChooser(shareIntent, "分享到"));
    }

    /**
     * 分享多张图片
     *
     * @param paths 路径的集合
     */
    public void shareMultipleImage(List<String> paths) {
        ArrayList<Uri> uriList = new ArrayList<>();
        for (String path : paths) {
            uriList.add(Uri.fromFile(new File(path)));
        }
        Intent shareIntent = new Intent();
        shareIntent.setAction(Intent.ACTION_SEND_MULTIPLE);
        shareIntent.putParcelableArrayListExtra(Intent.EXTRA_STREAM, uriList);
        shareIntent.setType("image/*");
        activity.startActivity(Intent.createChooser(shareIntent, "分享到"));
    }


    private void handleListener(int type, List<String> list, String title, String content) {
        if (listener != null) {
            listener.OnHandleOk(type, list, title, content);
        }
    }

    /**
     * 处理分享的文本
     *
     * @param intent
     */
    private void handleSendText(Intent intent) {
        String sharedText = intent.getStringExtra(Intent.EXTRA_TEXT);
        String sharedTitle = intent.getStringExtra(Intent.EXTRA_TITLE);
        handleListener(TYPE_TEXT, null, sharedTitle, sharedText);
    }

    /**
     * 处理分享的单张照片
     *
     * @param intent
     */
    private void handleSendImage(Intent intent) {
        Uri imageUri = intent.getParcelableExtra(Intent.EXTRA_STREAM);
        if (imageUri == null)
            return;
        List<String> list = new ArrayList<>();
        list.add(UriUtils.getRealPathFromURI(activity, imageUri));
        handleListener(TYPE_IMAGE, list, null, null);
    }

    /**
     * 处理分享的多张照片
     *
     * @param intent
     */
    private void handleSendMultipleImages(Intent intent) {

        ArrayList<Uri> imageUris = intent.getParcelableArrayListExtra(Intent.EXTRA_STREAM);
        // LUtils.e("golemscorner", imageUris.toString());
        if (imageUris == null)
            return;
        List<String> list = new ArrayList<>();
        for (Uri uri : imageUris) {
            if (uri != null)
                list.add(UriUtils.getRealPathFromURI(activity, uri));
        }
        handleListener(TYPE_IMAGE, list, null, null);
    }


    public interface OnShareDataOkListener {
        void OnHandleOk(int type, List<String> list, String title, String content);
    }
}
