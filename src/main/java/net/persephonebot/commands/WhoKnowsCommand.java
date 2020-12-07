package net.persephonebot.commands;

import com.jagrosh.jdautilities.command.CommandEvent;

public class WhoKnowsCommand extends BaseCommand {
    public WhoKnowsCommand() {
        this.name = "whoknows";
        this.aliases = new String[]{"wk", "npwk"};
        this.help = "Shows who knows a specific artist";
    }

    @Override
    protected void execute(CommandEvent event) {
        event.reply("Not implemented");
    }
}
