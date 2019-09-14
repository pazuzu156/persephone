<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class HomeController extends Controller
{
    private $cmds = [
        'About',
        'BandInfo',
        'YouTube',
    ];

    public function index()
    {
        return $this->_page('home');
    }

    public function help()
    {
        return $this->_page('help')->with('cmds', $this->cmds);
    }

    public function getDoc($doc)
    {
        $docpath = resource_path().'/views/docs/markdown/'.$doc.'.md';

        return $this->_page('docs.getdoc', "$doc Help")->with([
            'doc' => $docpath,
            'cmds' => $this->cmds,
        ]);
    }
}
