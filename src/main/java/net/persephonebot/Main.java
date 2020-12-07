package net.persephonebot;

import javax.security.auth.login.LoginException;

import com.jagrosh.jdautilities.command.CommandClientBuilder;

import net.dv8tion.jda.api.JDABuilder;
import net.dv8tion.jda.api.OnlineStatus;
import net.dv8tion.jda.api.entities.Activity;
import net.persephonebot.commands.AboutCommand;
import net.persephonebot.commands.RecentCommand;
import net.persephonebot.commands.WhoKnowsCommand;
import net.persephonebot.utils.Listener;

public class Main {
    public static String Version = "0.1";

    public static void main(String[] args) {
        // EventWaiter waiter = new EventWaiter();
        CommandClientBuilder client = new CommandClientBuilder()
            .useDefaultGame()
            .setOwnerId(BotConfig.cfg().ownerID)
            .setPrefix(BotConfig.cfg().prefix)
            .setStatus(OnlineStatus.DO_NOT_DISTURB)
            .setActivity(Activity.listening("Music"));

        client.addCommands(new AboutCommand(),
            new RecentCommand(),
            new WhoKnowsCommand());

        try {
            JDABuilder.createDefault(BotConfig.cfg().token)
                .addEventListeners(client.build(), new Listener())
                .build();
        } catch (LoginException e) {
            e.printStackTrace();
        }
    }
}
