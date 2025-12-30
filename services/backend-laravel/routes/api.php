<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes - Logistics Core
|--------------------------------------------------------------------------
*/

// Health check endpoint for Traefik and Frontend
Route::get('/core/status', function () {
    return response()->json([
        'status' => 'online',
        'service' => 'Laravel Logistics Core',
        'message' => 'System is ready to handle inventory data',
        'timestamp' => now()
    ]);
});