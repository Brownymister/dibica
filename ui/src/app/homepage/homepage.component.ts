import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.css']
})
export class HomepageComponent implements OnInit {

  constructor() { 
  }

  title = 'Generate your digital bithday card now!';
  
  ngOnInit(): void {
  }

}
