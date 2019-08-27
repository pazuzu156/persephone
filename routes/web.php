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

Route::get('/', function () {
    return view('welcome');
});

Route::name('auth.api.')->group(function ($route) {
    $route->prefix('auth/api')->group(function ($route) {
        $route->get('/failed', function () {
            return view('auth.api.failedToken');
        });
        $route->get('/expired', function () {
            return view('auth.api.expiredToken');
        });
        $route->get('/continue', 'Auth\LoginController@beginAuthFlow');
    });
});
