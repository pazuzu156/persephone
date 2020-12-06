package net.persephonebot.commands;

import java.text.DateFormat;
import java.text.SimpleDateFormat;
import java.util.Date;

import com.jagrosh.jdautilities.command.Command;
import com.jagrosh.jdautilities.command.CommandEvent;

import org.apache.commons.lang3.StringUtils;

import net.dv8tion.jda.api.EmbedBuilder;
import net.dv8tion.jda.api.entities.Member;
import net.dv8tion.jda.api.entities.Role;
import net.dv8tion.jda.api.entities.SelfUser;
import net.persephonebot.Main;
import net.persephonebot.utils.Strings;

public class AboutCommand extends Command {
    public AboutCommand() {
        this.name = "about";
        this.help = "Gets information about the bot";
    }

    @Override
    protected void execute(CommandEvent evt) {
        SelfUser user = evt.getJDA().getSelfUser();
        DateFormat sdf = new SimpleDateFormat("k:mm a z");
        Date date = new Date();
        EmbedBuilder eb = new EmbedBuilder()
            .setTitle("About Persephone", null)
            .setDescription("Persephone is abot written in Java. Version: "+Main.Version)
            .setColor(0x00ff00)
            .setThumbnail(user.getAvatarUrl())
            .addField("Name", Strings.User(user), false)
            .addField("ID", user.getId(), false)
            .addField("Roles", roles(evt), false)
            .addField("Source", "https://github.com/pazuzu156/persephone", true)
            .addField("Website", "https://persephonebot.net", true)
            .setFooter("Command invoked by: "+Strings.User(evt.getAuthor())+" \u2022 Today at "+sdf.format(date),
                evt.getAuthor().getAvatarUrl());

        evt.reply(eb.build());
    }

    private String roles(CommandEvent evt) {
        Member member = evt.getGuild().getMember(evt.getSelfUser());
        StringBuilder sb = new StringBuilder();

        for (Role role : member.getRoles()) {
            sb.append(role.getName()+", ");
        }

        return StringUtils.strip(sb.toString(), ", ");
    }
}
