<?php

namespace App\Http\Controllers;

use Auth;
use Illuminate\Http\Request;
use RestCord\DiscordClient;

class HomeController extends Controller
{
    private $cmds = [
        'About',
        'BandInfo',
        'YouTube',
    ];

    public function index()
    {
        return $this->_page('home')->with('user', $this->user());
    }

    public function help()
    {
        return $this->_page('help')->with([
            'cmds' => $this->cmds,
            'user' => $this->user(),
        ]);
    }

    public function getDoc($doc)
    {
        $docpath = resource_path().'/views/docs/markdown/'.$doc.'.md';

        return $this->_page('docs.getdoc', "$doc Help")->with([
            'doc' => $docpath,
            'cmds' => $this->cmds,
            'user' => $this->user(),
        ]);
    }

    private function user()
    {
        $user = '';

        if (Auth::check()) {
            $discord = new DiscordClient([
                'token' => Auth::user()->discord_token,
                'tokenType' => 'OAuth',
            ]);

            try {
                $user = $discord->user->getCurrentUser();
            } catch (\Exception $ex) {
                return redirect()->route('auth.reauthDiscord');
            }
        }

        return $user;
    }
}
