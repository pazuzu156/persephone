package net.persephonebot.database;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.HashMap;
import java.util.Map;

import net.dv8tion.jda.api.entities.User;

public class DBCrown extends Db {
    public DBCrown() throws SQLException {
        super();
    }

    /**
     * Gets a user's crowns.
     *
     * @param user
     * @return
     * @throws SQLException
     */
    public Map<Integer, String[]> getCrowns(User user, int limit, int offset) throws SQLException {
        ResultSet res = select(
            "crowns",
            "discord_id", "=", user.getId(),
            orderDesc("play_count"),
            limit(limit),
            offset(offset)
        );

        return get(res);
    }

    public Map<Integer, String[]> getCrowns(User user) throws SQLException {
        ResultSet res = select(
            "crowns",
            "discord_id", "=", user.getId()
        );

        return get(res);
    }

    private Map<Integer, String[]> get(ResultSet res) throws SQLException {
        Map<Integer, String[]> crowns = new HashMap<Integer, String[]>();
        int x = 0;

        while (res.next()) {
            crowns.put(x, new String[]{
                res.getString("artist"),
                res.getString("play_count")
            });
            x++;
        }

        close();

        return crowns;
    }
}
