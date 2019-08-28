<?php

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

use RestCord\DiscordClient;

Route::get('/', function () {
    return view('welcome');
});

Route::name('auth.')->prefix('auth')->group(function ($route) {
    $route->get('/authenticate/{discordId}/{token}', 'Auth\LoginController@authenticateUserWithToken');
    $route->get('/continue', 'Auth\LoginController@beginAuthFlow');

    $route->name('discord.')->prefix('/discord')->group(function ($route) {
        $route->get('/', function () {
            return Socialite::with('discord')->scopes(['identify', 'guilds'])->redirect();
        })->name('begin');
        $route->get('/callback', 'Auth\LoginController@discordCallback');
    });

    $route->name('lastfm.')->prefix('/lastfm')->group(function ($route) {
        $route->get('/callback/{discordId}', 'Auth\LoginController@lastfmCallback');
        $route->get('/{discordId}', 'Auth\LoginController@beginLastfmAuthentication')->name('begin');
    });

    $route->get('/complete', function () {
        return view('auth.complete');
    });
});

Route::name('auth.api.')->group(function ($route) {
    $route->prefix('auth/api')->group(function ($route) {
        $route->get('/failed', function () {
            return view('auth.api.failedToken');
        });
        $route->get('/expired', function () {
            return view('auth.api.expiredToken');
        });
    });
});
