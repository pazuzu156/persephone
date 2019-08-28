@extends('layout.main')

@section('content')
<p>
    Your request to begin the authentication process has been accepted. You're already authenticated with Discord, so we'll skip that
    and head right into <a href="{{ route('auth.lastfm.begin', compact('discordId')) }}">loggin into Last.fm</a>
</p>
@endsection
