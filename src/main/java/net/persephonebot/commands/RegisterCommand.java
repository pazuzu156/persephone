package net.persephonebot.commands;

import java.sql.SQLException;

import com.jagrosh.jdautilities.command.CommandEvent;

import net.persephonebot.database.DBUser;

public class RegisterCommand extends BaseCommand {
    public RegisterCommand() {
        this.name = "register";
        this.help = "Requests a registration URL";
    }

    @Override
    protected void execute(CommandEvent event) {
        try {
            DBUser dbu = new DBUser(event.getAuthor());
            if (dbu.exists()) {
                event.reply("You're already logged in");
            } else {
                event.reply(String.format("Please go to %s to begin registration", config.webLoginUrl.toString()));
            }
        } catch (SQLException e) {
            e.printStackTrace();
        }
    }
}
