package net.persephonebot;

import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;

import javax.security.auth.login.LoginException;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory;
import com.jagrosh.jdautilities.command.CommandClientBuilder;

import net.dv8tion.jda.api.JDABuilder;
import net.dv8tion.jda.api.OnlineStatus;
import net.dv8tion.jda.api.entities.Activity;
import net.persephonebot.commands.AboutCommand;
import net.persephonebot.utils.Config;
import net.persephonebot.utils.Listener;

public class Main {
    public static String Version = "0.1";

    public static void main(String[] args) {
        InputStream is = Main.class.getClassLoader().getResourceAsStream("config.yml");
        ObjectMapper mapper = new ObjectMapper(new YAMLFactory());
        Config cfg = null;

        try {
            cfg = mapper.readValue(new InputStreamReader(is), Config.class);
        } catch (IOException e) {
            e.printStackTrace();
        }

        // EventWaiter waiter = new EventWaiter();
        CommandClientBuilder client = new CommandClientBuilder()
            .useDefaultGame()
            .setOwnerId(cfg.ownerID)
            .setPrefix(",")
            .setStatus(OnlineStatus.DO_NOT_DISTURB)
            .setActivity(Activity.listening("Music"));

        client.addCommands(new AboutCommand());

        try {
			JDABuilder.createDefault(cfg.token)
			    .addEventListeners(client.build(), new Listener())
			    .build();
		} catch (LoginException e) {
			e.printStackTrace();
        }
    }
}
