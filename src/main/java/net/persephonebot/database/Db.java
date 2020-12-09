package net.persephonebot.database;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
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
     * Runs a select query.
     *
     * @param table - the table to query
     * @param options - query options
     * @return
     * @throws SQLException
     */
    public ResultSet select(String table, String... options) throws SQLException {
        String sql = "SELECT * FROM "+table;
        PreparedStatement stmt = null;

        if (options.length > 0) {
            if (options.length > 1) {
                sql += String.format(" WHERE `%s` %s ?", options[0], options[1]);
                if (options.length > 3 && options[3] != null) {
                    // sql += " "+options[3];
                    for (int i = 3; i < options.length; i++) {
                        sql += " "+options[i];
                    }
                }

                stmt = prepare(sql);
                stmt.setString(1, options[2]);
            }
        } else {
            stmt = prepare(sql);
        }

        return stmt.executeQuery();
    }

    /**
     * Writes an ascending order string.
     *
     * @param by - Item to order by
     * @return
     */
    public String orderAsc(String by) {
        return String.format("ORDER BY `%s` ASC", by);
    }

    /**
     * Writes a descending order string.
     *
     * @param by - Item to order by
     * @return
     */
    public String orderDesc(String by) {
        return String.format("ORDER BY `%s` DESC", by);
    }

    public String limit(int limit) {
        return String.format("LIMIT %d", limit);
    }

    public String offset(int offset) {
        return String.format("OFFSET %d", offset);
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
