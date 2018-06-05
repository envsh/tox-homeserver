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
import android.provider.MediaStore;

/**
 * CdLibsTest     com.march.libs.utils
 * Created by 陈栋 on 16/3/25.
 * 功能:
 */
public class UriUtils {

    /**
     * 从uri获取path
     *
     * @param uri
     * @return
     */
    public static String getRealPathFromURI(Context context, Uri uri) {
        if (null == uri) return null;
        final String scheme = uri.getScheme();
        String data = null;
        if (scheme == null)
            data = uri.getPath();
        else if (ContentResolver.SCHEME_FILE.equals(scheme)) {
            data = uri.getPath();
        } else if (ContentResolver.SCHEME_CONTENT.equals(scheme)) {
            Cursor cursor = context.getContentResolver().query(uri, new String[]{MediaStore.Images.ImageColumns.DATA}, null, null, null);
            if (null != cursor) {
                if (cursor.moveToFirst()) {
                    int index = cursor.getColumnIndex(MediaStore.Images.ImageColumns.DATA);
                    if (index > -1) {
                        data = cursor.getString(index);
                    }
                }
                cursor.close();
            }
        }
        return data;
    }
}
