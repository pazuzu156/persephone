package net.persephonebot.commands;

import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;

import com.jagrosh.jdautilities.command.CommandEvent;

import net.persephonebot.utils.Db;

public class RegisterCommand extends BaseCommand {
    public RegisterCommand() {
        this.name = "register";
        this.help = "Requests a registration URL";
    }

    @Override
    protected void execute(CommandEvent event) {
        try {
            Db db = new Db();

            if (db.isOpen()) {
                PreparedStatement stmt = db.prepare("SELECT * FROM `users` WHERE `discord_id` = ?");
                stmt.setLong(1, event.getAuthor().getIdLong());
                ResultSet res = stmt.executeQuery();

                if (res.first()) {
                    event.reply("You're already logged in");
                } else {
                    event.reply(String.format("Please go to %s to begin registration", config.webLoginUrl.toString()));
                }
            }
        } catch (SQLException e) {
            e.printStackTrace();
        }
    }
}
