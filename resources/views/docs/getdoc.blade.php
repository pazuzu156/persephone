@extends('layout.main')

@section('content')
@include('docs.nav')

<div class="col-lg-10 ml-auto cmd-div">
    @if(file_exists($doc) && filesize($doc) > 0)
    @php
        $handle = fopen($doc, 'r');
        $content = fread($handle, filesize($doc));
        fclose($handle);
    @endphp

    {!! Markdown::parse($content) !!}
    @else
    <p>Help information could not be found for the requested command!</p>
    @endif
</div>

@endsection
