<div class="col-lg-2">
    <h5 class="sidebar-head">General</h5>
    @if(strtolower($pageTitle) != 'help')
    <a href="{{ url('/help') }}">Help Home</a>
    @else
    Help Home
    </a>
    @endif
    <br>
    <br>
    <h5 class="sidebar-head">Commands</h5>
    @foreach($cmds as $cmd)
    @if(strtolower($pageTitle) == strtolower($cmd).' help')
    {{ $cmd }}<br>
    @else
    <a href="{{ url('/help/'.strtolower($cmd)) }}">{{ $cmd }}</a><br>
    @endif
    @endforeach
</div>
