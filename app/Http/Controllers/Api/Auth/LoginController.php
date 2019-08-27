<?php

namespace App\Http\Controllers\Api\Auth;

use App\Login;
use Illuminate\Http\Request;
use App\Http\Controllers\Controller;

class LoginController extends Controller
{
    public function createLoginRequestToken($discordId)
    {
        $token = str_random(16); // 16 character token
        $exists = Login::where('discord_id', '=', $discordId);

        if ($exists->count()) {
            $login = $exists->first();
        } else {
            $login = new Login();
            $login->discord_id = $discordId;
        }

        $login->request_token = $token;
        $login->expires = now()->addMinutes(6);

        if ($login->save()) {
            $response = [
                'request_token' => $token,
                'expires' => $login->expires->timestamp,
                'expires_string' => "expires in {$login->expires->diff()->i} minutes",
                'error' => false,
            ];
        } else {
            $response = [
                'error' => true,
                'message' => 'There was an issue attempting to retrieve a request token. Please run the login command again.'
            ];
        }

        return response($response)->header('Content-type', 'application/json');
    }

    public function authenticateUserWithToken($discordId, $token)
    {
        $login = Login::where('discord_id', '=', $discordId)->where('request_token', '=', $token);

        if ($login->count()) {
            $request = $login->first();
            $request->delete();

            if ($request->expires <= now()) {
                return redirect('/auth/api/expired');
            }

            return redirect('/auth/api/continue')->with('token', $token);
        }

        return redirect('/auth/api/failed');
    }
}
