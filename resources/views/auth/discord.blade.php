@extends('layout')

@section('content')
Your request to begin the authentication process has been accepted. Go ahead and start by <a href="{{ route('auth.discord.begin') }}">Logging in with Discord</a>
@endsection
