import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { RouterModule, Routes } from '@angular/router'

import { AppComponent } from './app.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { HomepageComponent } from './homepage/homepage.component';
import { TemplateCardComponent } from './template-card/template-card.component';
import { CreatePageComponent } from './create-page/create-page.component';
import { ViewPageComponent } from './view-page/view-page.component';
import { CollectionComponent } from './collection/collection.component';

const routes: Routes = [
  { path: 'home', component: HomepageComponent },
  { path: 'create', component: CreatePageComponent },
  { path: 'card', component: ViewPageComponent },
  { path: 'collection', component: CollectionComponent },
  { path: '', redirectTo: '/home', pathMatch: 'full' },
]

@NgModule({
  declarations: [
    AppComponent,
    HomepageComponent,
    TemplateCardComponent,
    CreatePageComponent,
    ViewPageComponent,
    CollectionComponent,
  ],
  imports: [
    BrowserModule,
    NgbModule,
    RouterModule.forRoot(routes)
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
