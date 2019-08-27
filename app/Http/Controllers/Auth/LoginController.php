<?php

namespace App\Http\Controllers\Auth;

use Illuminate\Http\Request;
use App\Http\Controllers\Controller;

class LoginController extends Controller
{
    public function beginAuthFlow()
    {
        dd(session()->get('token'));
    }
}
