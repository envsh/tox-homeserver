package im.fixlan.golem;


import java.lang.*;

public class PendingIntents {
    public static int PendingCount = 0;
    public static int GetPendingCount() {
        return PendingIntents.PendingCount;
    }

    public static String PendingData = ""; // TODO json data, or simple seperate by a string
    public static String GetPendingData() {
        String data = PendingIntents.PendingData;
        PendingIntents.PendingData = "";
        PendingIntents.PendingCount = 0;
        return data;
    }

    public static void AddItem(String mime, String data) {
        PendingIntents.PendingCount +=1;
        String sep = "-:::-";
        PendingIntents.PendingData += sep + mime + sep + data;
    }
}
