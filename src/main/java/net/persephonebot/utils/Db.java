package net.persephonebot.utils;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.PreparedStatement;
import java.sql.SQLException;
import java.util.Map;
import java.util.Properties;

import net.persephonebot.BotConfig;

public class Db {
    private Connection _connection;

    public Db() throws SQLException {
        open();
    }

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

    public Boolean isOpen() {
        return (_connection == null) ? false : true;
    }

    public PreparedStatement prepare(String sql) throws SQLException {
        return _connection.prepareStatement(sql);
    }

    public Db close() throws SQLException {
        _connection.close();

        return this;
    }
}
