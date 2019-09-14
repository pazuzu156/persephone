<?php

namespace App\Http\Controllers;

use Illuminate\Foundation\Bus\DispatchesJobs;
use Illuminate\Routing\Controller as BaseController;
use Illuminate\Foundation\Validation\ValidatesRequests;
use Illuminate\Foundation\Auth\Access\AuthorizesRequests;

class Controller extends BaseController
{
    use AuthorizesRequests, DispatchesJobs, ValidatesRequests;

    protected function _page($view, $title = '')
    {
        if ($title == '') {
            $r = explode('.', $view);
            $title = $r[count($r)-1];
        }

        $title = ucwords(str_replace('-', ' ', $title));

        return view($view)->with('pageTitle', $title);
    }
}
