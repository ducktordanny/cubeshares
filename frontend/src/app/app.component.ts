import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'cubeshares-root',
  templateUrl: './app.template.html',
  styleUrl: './app.style.scss',
  imports: [RouterOutlet],
})
export class AppComponent {}
