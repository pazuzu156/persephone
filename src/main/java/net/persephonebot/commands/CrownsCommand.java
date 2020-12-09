package net.persephonebot.commands;

import java.sql.SQLException;
import java.time.Instant;
import java.util.Map;

import com.jagrosh.jdautilities.command.CommandEvent;

import org.apache.commons.lang3.StringUtils;
import org.apache.commons.lang3.math.NumberUtils;

import net.dv8tion.jda.api.EmbedBuilder;
import net.dv8tion.jda.api.entities.Member;
import net.dv8tion.jda.api.entities.User;
import net.dv8tion.jda.api.requests.RestAction;
import net.persephonebot.database.DBUser;

public class CrownsCommand extends BaseCommand {
    User user;
    int page;

    public CrownsCommand() {
        this.name = "crowns";
        this.help = "Shows a list of crowns for yourself or a given user";
        this.arguments = "[user] [page]";
    }

    @Override
    protected void execute(CommandEvent event) {
        user = event.getAuthor(); // set author as user for default
        page = 1;

        if (event.getArgs().length() > 0) {
            String[] args = event.getArgs().split(" ");

            for (String arg : args) {
                if (NumberUtils.isNumber(arg)) {
                    page = NumberUtils.toInt(arg);
                } else {
                    RestAction<Member> rem = event.getGuild().retrieveMemberById(StringUtils.strip(arg, "<@!>"));
                    Member m = rem.complete();
                    user = m.getUser();
                }
            }
        }

        try {
            int maxPerPage = 10;
            int offset = 0;
            DBUser dbu = new DBUser(user);
            Map<Integer, String[]> crowns = dbu.getCrowns();
            int count = crowns.size();
            int pages = (int) Math.ceil(Float.valueOf(count) / Float.valueOf(maxPerPage));
            StringBuilder sb = new StringBuilder();

            EmbedBuilder eb = new EmbedBuilder()
                .setTitle(String.format("%d crowns for %s", crowns.size(), user.getName()))
                .setColor(randomColor())
                .setTimestamp(Instant.now());

            if (page <= pages) {
                if (page > 1) {
                    offset = (page - 1) * maxPerPage;
                }
                dbu = new DBUser(user);
                crowns = dbu.getCrowns(maxPerPage, offset);

                for (int i = 0; i < crowns.size(); i++) {
                    sb.append(String.format(
                        "%d. :crown: %s with %s plays\n",
                        (i+1)+offset,
                        crowns.get(i)[0],
                        crowns.get(i)[1]
                    ));
                }

                eb.setDescription(StringUtils.strip(sb.toString(), "\n")).setFooter(
                    footerText(event)+String.format(" | Page %d/%d", page, pages),
                    event.getAuthor().getAvatarUrl()
                );
                event.reply(eb.build());
            }



        } catch (SQLException e) {
            e.printStackTrace();
        }
    }
}
