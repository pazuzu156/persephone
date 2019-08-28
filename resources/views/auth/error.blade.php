@extends('layout.main')

@section('content')

@if(isset($reason))
<p>An error occurred when attempting to access this page. Reason: {{ $reason }}</p>
@else
<p>An error occurred when attempting to login with Discord. Error: {{ $request->error_description }}</p>
@endif

@endsection
