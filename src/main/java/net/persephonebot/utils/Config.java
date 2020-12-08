package net.persephonebot.utils;

import java.net.URL;

public class Config {
    public String token;
    public String ownerID;
    public String prefix;
    public LastFM lastfm;
    public Database database;
    public Youtube youtube;
    public URL webLoginUrl;

    public class LastFM {
        public String apikey;
        public String secret;
    }

    public class Database {
        public String hostname;
        public int port;
        public String username;
        public String password;
        public String name;
    }
    public class Youtube {
        public String apikey;
    }
}
