package net.persephonebot.commands;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.Properties;

import com.jagrosh.jdautilities.command.CommandEvent;

import de.umass.lastfm.PaginatedResult;
import de.umass.lastfm.Track;
import de.umass.lastfm.User;
import net.dv8tion.jda.api.EmbedBuilder;

public class RecentCommand extends BaseCommand {
    public RecentCommand() {
        this.name = "recent";
        this.help = "Shows a list of recent tracks";
    }

    @Override
    protected void execute(CommandEvent event) {
        EmbedBuilder eb = new EmbedBuilder()
                .setTitle("Recent Tracks")
                .setThumbnail(event.getAuthor().getAvatarUrl())
                .setFooter(footerText(event), event.getAuthor().getAvatarUrl())
                .setColor(randomColor());

        Connection conn = null;
        try {
            Properties props = new Properties();
            props.setProperty("passwordCharacterEncoding", "UTF-8");
            conn = DriverManager
                    .getConnection(String.format("jdbc:mariadb://%s:%d/%s?user=%s",
                        config.database.hostname,
                        config.database.port,
                        config.database.name,
                        config.database.username,
                        config.database.password), props);
        } catch (SQLException e) {
            e.printStackTrace();
        }

        if (conn != null) {
            try {
                PreparedStatement stmt = conn.prepareStatement("SELECT * FROM `users` WHERE `discord_id` = ?");
                stmt.setLong(1, event.getAuthor().getIdLong());
                ResultSet res = stmt.executeQuery();
                String lastfm = "";

                while (res.next()) {
                    lastfm = res.getString("lastfm");
                }
                conn.close();

                PaginatedResult<Track> recentTracks = User.getRecentTracks(lastfm, 1, 5, config.lastfm.apikey);
                StringBuilder sb = new StringBuilder();

                for (Track track : recentTracks) {
                    if (track.isNowPlaying()) {
                        eb.addField("Currently Playing", track.getArtist()+" - "+track.getName(), false);
                    } else {
                        sb.append(String.format("%s - %s\n", track.getArtist(), track.getName()));
                    }
                }

                eb.addField("Previous Tracks", sb.toString(), false);
            } catch (SQLException e) {
                e.printStackTrace();
            }


            event.reply(eb.build());
        }


    }
}
