package net.persephonebot.commands;

import java.awt.Color;

import com.jagrosh.jdautilities.command.Command;
import com.jagrosh.jdautilities.command.CommandEvent;

import net.persephonebot.BotConfig;
import net.persephonebot.utils.Config;
import net.persephonebot.utils.Strings;

public abstract class BaseCommand extends Command {
    protected Config config = null;

    public BaseCommand() {
        config = BotConfig.cfg();
        this.guildOnly = true;
    }

    /**
     * Writes common footer text.
     *
     * @param event
     * @return footer test string
     */
    public String footerText(CommandEvent event) {
        return "Command invoked by: " + Strings.User(event.getAuthor());
    }

    /**
     * Generates a random color. Used for embeds.
     *
     * @return Random color
     */
    public Color randomColor() {
        return new Color((int) (Math.random() * 0x1000000));
    }
}
