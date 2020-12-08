package net.persephonebot.commands;

import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;

import com.jagrosh.jdautilities.command.CommandEvent;

import de.umass.lastfm.ImageSize;
import de.umass.lastfm.PaginatedResult;
import de.umass.lastfm.Track;
import de.umass.lastfm.User;
import net.dv8tion.jda.api.EmbedBuilder;
import net.persephonebot.utils.Db;

public class RecentCommand extends BaseCommand {
    public RecentCommand() {
        this.name = "recent";
        this.help = "Shows a list of recent tracks";
    }

    @Override
    protected void execute(CommandEvent event) {
        EmbedBuilder eb = new EmbedBuilder()
            .setTitle("Recent Tracks")
            .setFooter(footerText(event), event.getAuthor().getAvatarUrl())
            .setColor(randomColor());

        try {
            Db db = new Db();

            if (db.isOpen()) {
                PreparedStatement stmt = db.prepare("SELECT * FROM `users` WHERE `discord_id` = ?");
                stmt.setLong(1, event.getAuthor().getIdLong());
                ResultSet res = stmt.executeQuery();
                String lastfm = "";
                String imageUrl = event.getAuthor().getAvatarUrl();

                while (res.next()) {
                    lastfm = res.getString("lastfm");
                }
                db.close();

                PaginatedResult<Track> recent = User.getRecentTracks(
                    lastfm, 1, 5, config.lastfm.apikey
                );
                StringBuilder sb = new StringBuilder();

                for (Track track : recent) {
                    if (track.isNowPlaying()) {
                        eb.addField(
                            "Currently Playing",
                            track.getArtist()+" - "+track.getName(),
                            false
                        );
                        imageUrl = track.getImageURL(ImageSize.EXTRALARGE);

                        if (imageUrl == null) {
                            track.getImageURL(ImageSize.LARGE);
                        }
                    } else {
                        sb.append(String.format("%s - %s\n",
                            track.getArtist(), track.getName()));
                    }
                }

                System.out.println(imageUrl);

                eb.addField("Previous Tracks", sb.toString(), false);
                eb.setThumbnail(imageUrl);
            }

            event.reply(eb.build());
        } catch (SQLException e) {
            e.printStackTrace();
        }
    }
}
