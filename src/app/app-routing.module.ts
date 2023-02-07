import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CommonModule } from '@angular/common';
import { PageComponent } from './page/page.component';
import { Page2Component } from './page2/page2.component'; 

const routes: Routes = [
    { path: '', component: PageComponent },
    { path: 'page2', component: Page2Component }
    ];

    @NgModule({
        declarations: [],
        imports: [
        CommonModule,
        RouterModule.forRoot(routes)
        ],
        exports: [RouterModule]
        })
export class AppRoutingModule { }
