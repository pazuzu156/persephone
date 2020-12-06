package net.persephonebot.utils;

import net.dv8tion.jda.api.entities.User;

public class Strings {
    /**
     * Return a string representation of a user.
     * @param user - The user to get the string for
     * @return User as a string: User#0000
     */
    public static String User(User user) {
        return user.getName() + "#" + user.getDiscriminator();
    }
}
