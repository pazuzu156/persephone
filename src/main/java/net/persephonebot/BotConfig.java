package net.persephonebot;

import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory;

import net.persephonebot.utils.Config;

public class BotConfig {
    public static Config cfg() {
        InputStream is = BotConfig.class.getClassLoader().getResourceAsStream("config.yml");
        ObjectMapper mapper = new ObjectMapper(new YAMLFactory());
        Config cfg = null;

        try {
            cfg = mapper.readValue(new InputStreamReader(is), Config.class);
        } catch (IOException e) {
            e.printStackTrace();
        }

        return cfg;
    }
}
