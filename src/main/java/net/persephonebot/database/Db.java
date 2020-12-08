package net.persephonebot.database;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.PreparedStatement;
import java.sql.SQLException;
import java.util.Properties;

import net.persephonebot.BotConfig;

public class Db {
    private Connection _connection;

    /**
     * Bot's Database class.
     *
     * @throws SQLException
     */
    public Db() throws SQLException {
        open();
    }

    /**
     * Opens a new database instance.
     *
     * @return
     * @throws SQLException
     */
    public Db open() throws SQLException {
        Properties props = new Properties();
        props.put("user", BotConfig.cfg().database.username);

        if (BotConfig.cfg().database.password != null) {
            props.put("password", BotConfig.cfg().database.password);
        }

        _connection = DriverManager.getConnection(
            String.format(
                "jdbc:mariadb://%s:%s/%s",
                BotConfig.cfg().database.hostname,
                BotConfig.cfg().database.port,
                BotConfig.cfg().database.name
            ), props
        );

        return this;
    }

    /**
     * Checks if a database connection is currently open.
     *
     * @return
     */
    public Boolean isOpen() {
        return (_connection == null) ? false : true;
    }

    /**
     * Used to prepare a sql statement.
     *
     * @param sql
     * @return
     * @throws SQLException
     */
    public PreparedStatement prepare(String sql) throws SQLException {
        return _connection.prepareStatement(sql);
    }

    /**
     * Closes database connection.
     *
     * @return
     * @throws SQLException
     */
    public Db close() throws SQLException {
        _connection.close();

        return this;
    }
}
