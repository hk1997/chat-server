package com.example.chatserver.firebase

import android.util.Log
import android.widget.Toast
import com.google.android.gms.tasks.OnCompleteListener
import com.google.firebase.messaging.FirebaseMessaging
import com.google.firebase.messaging.FirebaseMessagingService

class ChatServerFirebaseMessagingService : FirebaseMessagingService() {

    fun generateFirebaseToken(){
        FirebaseMessaging.getInstance().token
            .addOnCompleteListener(OnCompleteListener { task ->
                if (!task.isSuccessful) {
                    Log.w(TAG, "getInstanceId failed", task.exception)
                    return@OnCompleteListener
                }

                // Get new Instance ID token
                val token = task.result

                // Log and toast
                Log.d(TAG, token)
            })

    }


    companion object{
        @JvmStatic  val TAG = "ChatServerFirebaseMessagingService"
    }

}