import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { LandingPageComponent } from './landing-page.component';

describe('LandingPageComponent', () => {
  let component: LandingPageComponent;
  let fixture: ComponentFixture<LandingPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        RouterTestingModule
      ],
      declarations: [ LandingPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LandingPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create "Landing Page"', async() => {
    expect(component).toBeTruthy();
  });
  it('should have title: "AttackOnCollege"', async() => {
   
    expect(component.title).toEqual('AttackOnCollege');
  });
  it('button should contain: "Login"', async() => {
  
   
    const btn = fixture.debugElement.nativeElement.querySelector('#login');
    expect(btn.innerHTML).toBe('Login');
  });
  it('button should contain: "Login"', async() => {
  
   
    const btn = fixture.debugElement.nativeElement.querySelector('#loginwhite');
    expect(btn.innerHTML).toBe('Login');
  });
});
