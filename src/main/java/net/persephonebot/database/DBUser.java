package net.persephonebot.database;

import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;

import net.dv8tion.jda.api.entities.User;

public class DBUser extends Db {
    private User _user;

    /**
     * Gets a user from the database.
     *
     * @param user
     * @throws SQLException
     */
    public DBUser(User user) throws SQLException {
        super();

        this._user = user;
    }

    /**
     * Returns first database result.
     *
     * @return
     * @throws SQLException
     */
    public ResultSet first() throws SQLException {
        PreparedStatement stmt = prepare("SELECT * FROM `users` WHERE `discord_id` = ?");
        stmt.setLong(1, this._user.getIdLong());
        ResultSet res = stmt.executeQuery();

        if (res.first()) {
            return res;
        }

        return null;
    }

    /**
     * Checks if requested user exists.
     *
     * @return
     * @throws SQLException
     */
    public Boolean exists() throws SQLException {
        return (first() != null) ? true : false;
    }

    /**
     * Just returns the JDA user instance.
     *
     * @return
     */
    public User jdaUser() {
        return this._user;
    }
}
