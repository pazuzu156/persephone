package net.persephonebot.commands;

import java.time.Instant;

import com.jagrosh.jdautilities.command.CommandEvent;

import org.apache.commons.lang3.StringUtils;

import net.dv8tion.jda.api.EmbedBuilder;
import net.dv8tion.jda.api.entities.Member;
import net.dv8tion.jda.api.entities.Role;
import net.dv8tion.jda.api.entities.SelfUser;
import net.persephonebot.Main;
import net.persephonebot.utils.Strings;

public class AboutCommand extends BaseCommand {
    public AboutCommand() {
        this.name = "about";
        this.help = "Gets information about the bot";
    }

    @Override
    protected void execute(CommandEvent event) {
        SelfUser user = event.getJDA().getSelfUser();
        EmbedBuilder eb = new EmbedBuilder()
            .setTitle("About Persephone", null)
            .setDescription("Persephone is abot written in Java. Version: "+Main.Version)
            .setColor(randomColor())
            .setThumbnail(user.getAvatarUrl())
            .addField("Name", Strings.User(user), false)
            .addField("ID", user.getId(), false)
            .addField("Roles", roles(event), false)
            .addField("Source", "https://github.com/pazuzu156/persephone", true)
            .addField("Website", "https://persephonebot.net", true)
            .setFooter(footerText(event), event.getAuthor().getAvatarUrl())
            .setTimestamp(Instant.now());

        event.reply(eb.build());
    }

    private String roles(CommandEvent event) {
        Member member = event.getGuild().getMember(event.getSelfUser());
        StringBuilder sb = new StringBuilder();

        for (Role role : member.getRoles()) {
            sb.append(role.getName()+", ");
        }

        return StringUtils.strip(sb.toString(), ", ");
    }
}
