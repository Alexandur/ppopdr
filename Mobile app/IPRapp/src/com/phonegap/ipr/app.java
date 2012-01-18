package com.phonegap.ipr;

import ipr.app.R;
import ipr.app.R.layout;
import android.os.Bundle;
import com.phonegap.*;

public class app extends DroidGap {
    /** Called when the activity is first created. */
    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        super.loadUrl("file:///android_asset/www/index.html");
    }
    
    @Override
    public void onBackPressed() {
	    if(appView.canGoBack()){
	       appView.goBack();
	    }
    }

}