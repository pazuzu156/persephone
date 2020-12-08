package net.persephonebot.commands;

import java.sql.SQLException;

import com.jagrosh.jdautilities.command.CommandEvent;

import de.umass.lastfm.ImageSize;
import de.umass.lastfm.PaginatedResult;
import de.umass.lastfm.Track;
import de.umass.lastfm.User;
import net.dv8tion.jda.api.EmbedBuilder;
import net.persephonebot.database.DBUser;

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
            DBUser dbu = new DBUser(event.getAuthor());
            if (dbu.exists()) {
                String lastfm = dbu.first().getString("lastfm");
                PaginatedResult<Track> recent = User.getRecentTracks(
                    lastfm,
                    1, 5,
                    config.lastfm.apikey
                );

                String imageUrl = dbu.jdaUser().getAvatarUrl();
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
                            imageUrl = track.getImageURL(ImageSize.LARGE);
                        }
                    } else {
                        sb.append(String.format(
                            "%s - %s\n",
                            track.getArtist(),
                            track.getName()
                        ));
                    }
                }

                eb.addField("Previous Tracks", sb.toString(), false);
                eb.setThumbnail(imageUrl);
            }

            event.reply(eb.build());
        } catch (SQLException e) {
            e.printStackTrace();
        }
    }
}
