<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Login extends Model
{
    protected $fillable = ['discord_id', 'request_token', 'expires'];
    protected $dates = ['expires'];
    protected $table = 'login';
}
