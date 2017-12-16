package io.dnesth.tofia;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;

import io.gomatcha.bridge.GoValue;
import io.gomatcha.matcha.MatchaView;

import android.view.KeyEvent;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        getSupportActionBar().hide();

        GoValue rootView = GoValue.withFunc("tox-homeserver/gofia New").call("")[0];
        setContentView(new MatchaView(this, rootView));
    }

    @Override
    public void onBackPressed() {
        //code......
        System.out.println("back pressed");
        GoValue ret = GoValue.withFunc("tox-homeserver/gofia OnBackPressed").call("")[0];
        if (ret.toLong()==1){
            // this.close();
            // MainActivity.exit();
            this.finish();
        }
    }


//    @Override
//    public boolean onKeyDown(int keyCode, KeyEvent event) {
//        GoValue ret = GoValue.withFunc("tox-homeserver/gofia OnKeyDown").call( Integer.toString((keyCode)), GoValue.WithInt(keyCode))[0];
//        return ret.toBool();
//    }
//
//    @Override
//    public boolean onKeyUp(int keyCode, KeyEvent event){
//        GoValue ret = GoValue.withFunc("tox-homeserver/gofia OnKeyUp").call( Integer.toString((keyCode)), GoValue.WithInt(keyCode))[0];
//        return ret.toBool();
//    }
//
//    @Override
//    public  boolean onKeyLongPress(int keyCode, KeyEvent event) {
//        GoValue ret = GoValue.withFunc("tox-homeserver/gofia OnKeyLongPress").call( Integer.toString((keyCode)), GoValue.WithInt(keyCode))[0];
//        return ret.toBool();
//    }
//
//    @Override
//    public  boolean onKeyMultiple(int keyCode, int count, KeyEvent event){
//        GoValue ret = GoValue.withFunc("tox-homeserver/gofia OnKeyMultiple").call( Integer.toString((keyCode)), GoValue.WithInt(keyCode))[0];
//        return ret.toBool();
//    }


}
