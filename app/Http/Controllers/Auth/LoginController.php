<?php

namespace App\Http\Controllers\Auth;

use Curl\Curl;
use Illuminate\Http\Request;
use App\Login;
use App\User;
use App\Http\Controllers\Controller;
use Socialite;
use RestCord\DiscordClient;

class LoginController extends Controller
{
    protected $_lfmBaseUri = 'https://www.last.fm/api/auth/';
    protected $_lfmBaseApiUri = 'https://ws.audioscrobbler.com/2.0/';

    public function beginAuthFlow()
    {
        // dd(session()->all());
        if (session()->has('token')) {
            // return view('auth.discord');
            $discordId = session('discordId');
            $user = User::where('discord_id', '=', $discordId);

            if ($user->count()) {
                return view('auth.lastfm')->with(compact('discordId'));
            }

            return view('auth.discord');
        }

        return view('auth.error')->with(['reason' => 'No or invalid request token used']);
    }

    public function authenticateUserWithToken($discordId, $token)
    {
        $login = Login::where('discord_id', '=', $discordId)->where('request_token', '=', $token);

        if ($login->count()) {
            $request = $login->first();

            if ($request->expires <= now()) {
                return redirect('/auth/api/expired');
            }

            return redirect('/auth/continue')->with(compact('token', 'discordId'));
        }

        return redirect('/auth/api/failed');
    }

    public function discordCallback(Request $request)
    {
        if (isset($request->error) && $request->error == 'access_denied') {
            return view('auth.error')->with(compact('request'));
        }

        if (isset($request->code)) {
            try {
                $r = Socialite::with('discord')->user();
            } catch (\Exception $ex) {
                return view('auth.error')->with('reason', 'An internal server error ocurred. You\'ll have to try the process again');
            }

            $discord = new DiscordClient([
                'token' => $r->token,
                'tokenType' => 'OAuth',
            ]);

            foreach ($discord->user->getCurrentUserGuilds() as $guild) {
                if ($guild->id == (int)env('DISCORD_GUILD_ID')) {
                    $user = new User();
                    $user->username = $r->name;
                    $user->email = $r->email;
                    $user->discord_id = $r->id;
                    $user->discord_token = $r->token;

                    if ($user->save()) {
                        return redirect()->route('auth.lastfm.begin', ['discordId' => $user->discord_id]);
                    }

                    return view('auth.error')->with('reason', 'There was an issue saving your Discord user data. Please try again later.');
                }
            }

            return view('auth.error')->with('reason', 'You are not a part of the Untrodden Corridors of Hades Discord server');
        }

        return view('auth.error')->with('reason', 'Unable to retrieve a Discord authentication token');
    }

    public function beginLastfmAuthentication($discordId)
    {
        return redirect($this->_lfmBaseUri.'?api_key='.env('LASTFM_KEY').'&cb='.env('LASTFM_REDIRECT_URI')."/{$discordId}");
    }

    public function lastfmCallback(Request $request, $discordId)
    {
        $user = User::where('discord_id', '=', $discordId);

        if ($user->count()) {
            $user = $user->first();
            $c = new Curl($this->_lfmBaseApiUri);
            $c->get('?method=auth.getSession&format=json&api_key='.env('LASTFM_KEY').'&token='.$request->token.'&api_sig='.$this->genApiSignature($request->token));

            if (isset($c->response->session)) {
                $user->lastfm = $c->response->session->name;
                $user->lastfm_token = $c->response->session->key;

                if ($user->save()) {
                    $login = Login::where('discord_id', '=', $discordId);
                    $login->delete();

                    return redirect('/auth/complete');
                }

                return view('auth.error')->with('reason', 'There was an issue saving your lastfm data. Please try again later.');
            }

            return view('auth.error')->with('reason', 'Last.fm Error: '.$c->response->message);
        }

        return view('auth.error')->with('reason', 'You do not have access to this page!');
    }

    private function genApiSignature($token)
    {
        return md5('api_key'.env('LASTFM_KEY').'methodauth.getSessiontoken'.$token.env('LASTFM_SECRET'));
    }
}
