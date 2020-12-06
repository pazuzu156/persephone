package net.persephonebot.utils;

import net.dv8tion.jda.api.events.ShutdownEvent;
import net.dv8tion.jda.api.hooks.ListenerAdapter;

public class Listener extends ListenerAdapter {
    @Override
    public void onShutdown(ShutdownEvent evt) {
        evt.getJDA().shutdown();
    }
}
