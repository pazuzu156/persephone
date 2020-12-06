package net.persephonebot;

import java.io.InputStream;
import java.io.InputStreamReader;

import javax.security.auth.login.LoginException;

import com.google.gson.Gson;
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
        InputStream is = Main.class.getClassLoader().getResourceAsStream("config.json");
        Gson gson = new Gson();
        Config cfg = gson.fromJson(new InputStreamReader(is), Config.class);

        // EventWaiter waiter = new EventWaiter();
        CommandClientBuilder client = new CommandClientBuilder()
            .useDefaultGame()
            .setOwnerId(cfg.getOwnerID())
            .setPrefix(",")
            .setStatus(OnlineStatus.DO_NOT_DISTURB)
            .setActivity(Activity.listening("Music"));

        client.addCommands(new AboutCommand());

        try {
			JDABuilder.createDefault(cfg.getToken())
			    .addEventListeners(client.build(), new Listener())
			    .build();
		} catch (LoginException e) {
			e.printStackTrace();
        }
    }
}
