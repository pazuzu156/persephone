<?php

namespace App\Http\Controllers\Api\Auth;

use App\Login;
use App\User;
use Illuminate\Http\Request;
use App\Http\Controllers\Controller;

class LoginController extends Controller
{
    public function createLoginRequestToken($discordId)
    {
        $user = User::where('discord_id', '=', $discordId);

        if ($user->count()) {
            $user = $user->first();

            if (!empty($user->lastfm)) {
                $response = [
                    'error' => true,
                    'message' => 'You\'re already logged in',
                ];

                return response($response)->header('Content-type', 'application/json');
            }
        }

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
}
