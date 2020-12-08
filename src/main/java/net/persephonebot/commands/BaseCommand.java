package net.persephonebot.commands;

import java.awt.Color;
import java.text.DateFormat;
import java.text.SimpleDateFormat;
import java.util.Date;

import com.jagrosh.jdautilities.command.Command;
import com.jagrosh.jdautilities.command.CommandEvent;

import net.persephonebot.BotConfig;
import net.persephonebot.utils.Config;
import net.persephonebot.utils.Strings;

public abstract class BaseCommand extends Command {
    protected Config config = null;

    public BaseCommand() {
        config = BotConfig.cfg();
    }

    public String footerText(CommandEvent event) {
        DateFormat sdf = new SimpleDateFormat("k:mm a z");
        Date date = new Date();

        return "Command invoked by: " + Strings.User(event.getAuthor()) + " \u2022 Today at " + sdf.format(date);
    }

    public Color randomColor() {
        return new Color((int) (Math.random() * 0x1000000));
    }
}
