package im.fixlan.golem;

import org.qtproject.qt5.android.bindings.QtActivity;
import android.os.*;
import android.content.*;
import android.app.*;
import android.Manifest;
import android.view.*;

import java.lang.String;
import android.content.Intent;
import java.io.File;
import android.net.Uri;
import android.util.Log;
import android.content.ContentResolver;
import android.webkit.MimeTypeMap;

import java.util.*;



public class GolemActivity extends QtActivity
{
    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        Log.d("golemscorner", "onCreate GolemActivity");
        getAndCacheIntentData(getIntent());
        handleSendText(getIntent());

        GolemActivity.requestPermission(this);
        // VideoRecorder.startRec(0);
    }

    // if we are opened from other apps:
    @Override
    public void onNewIntent(Intent intent) {
        Log.d("golemscorner", "onNewIntent");
        super.onNewIntent(intent);

        // processIntent();
        getAndCacheIntentData(intent);
        handleSendText(intent);
    }

    // 好像得到的消息有问题，有时会得到旧的分享消息？
    private void getAndCacheIntentData(Intent intent) {
        //如果你按照前面的要求注册了Activity，你可以在Activity中使用下面的代码处理分享得到的数据
        ShareSysUtils shareSysUtils = new ShareSysUtils(intent);
        shareSysUtils.handleShare(new ShareSysUtils.OnShareDataOkListener() {
                @Override
                public void OnHandleOk(int type, List<String> list, String title, String content) {
                    if (type == ShareSysUtils.TYPE_TEXT) {
                        Log.e("golemscorner", "分享文本是 " + title + "   " + content);
                        PendingIntents.AddItem("text", title + "   " + content);
                    } else if (type == ShareSysUtils.TYPE_IMAGE) {
                        Log.e("golemscorner", "分享图片是 " + list.toString());
                        PendingIntents.AddItem("image", list.toString());
                    }
                }
            });
    }

    //一行代码发起分享，会调起QQ,微信等App
    public void ShareImg(View view) {
        ShareSysUtils shareSysUtils = new ShareSysUtils(this);
        shareSysUtils.shareSingleImage("我是图片路径path");
    }

    public void ShareTxt(View view) {
        ShareSysUtils shareSysUtils = new ShareSysUtils(this);
        shareSysUtils.shareText("我是title", "我是content");
    }

    ///
    /**
     * 处理分享的文本
     *
     * @param intent
     */
    private void handleSendText(Intent intent) {
        String sharedText = intent.getStringExtra(Intent.EXTRA_TEXT);
        String sharedTitle = intent.getStringExtra(Intent.EXTRA_TITLE);
        // handleListener(TYPE_TEXT, null, sharedTitle, sharedText);
        Log.d("golemscorner, intent text:", sharedTitle + "   " + sharedText);
    }

    ///
    private void processIntent(){
      Intent intent = getIntent();

      Uri intentUri;
      String intentScheme;
      String intentAction;
      // we are listening to android.intent.action.SEND
      if (intent.getAction().equals("android.intent.action.SEND")){
             intentAction = "SEND";
              Bundle bundle = intent.getExtras();
              intentUri = (Uri)bundle.get(Intent.EXTRA_STREAM);
      } else {
              Log.d("golemscorner Intent unknown action:", intent.getAction());
              return;
      }
      Log.d("golemscorner action:", intentAction);
      if (intentUri == null){
            Log.d("golemscorner Intent URI:", "is null");
            return;
      }

      Log.d("golemscorner Intent URI:", intentUri.toString());

      // content or file
      intentScheme = intentUri.getScheme();
      if (intentScheme == null){
            Log.d("golemscorner Intent URI Scheme:", "is null");
            return;
      }
      if(intentScheme.equals("file")){
            // URI as encoded string
            Log.d("golemscorner Intent File URI: ", intentUri.toString());
            // setFileUrlReceived(intentUri.toString());
            // we are done Qt can deal with file scheme
            return;
      }
      if(!intentScheme.equals("content")){
              Log.d("golemscorner Intent URI unknown scheme: ", intentScheme);
              return;
      }
      // ok - it's a content scheme URI
      // we will try to resolve the Path to a File URI
      // if this won't work or if the File cannot be opened,
      // we'll try to copy the file into our App working dir via InputStream
      // hopefully in most cases PathResolver will give a path

      // you need the file extension, MimeType or Name from ContentResolver ?
      // here's HowTo get it:
      Log.d("golemscorner Intent Content URI: ", intentUri.toString());
      ContentResolver cR = this.getContentResolver();
      MimeTypeMap mime = MimeTypeMap.getSingleton();
      String fileExtension = mime.getExtensionFromMimeType(cR.getType(intentUri));
      Log.d("golemscorner Intent extension: ",fileExtension);
      String mimeType = cR.getType(intentUri);
      Log.d("golemscorner Intent MimeType: ",mimeType);
    }

    public static void requestPermission(QtActivity activity){
        activity.requestPermissions(new String[]{Manifest.permission.WRITE_EXTERNAL_STORAGE,
                                                 Manifest.permission.CAMERA}, 1);
    }

}
