import { Component, Input, model, OnInit } from '@angular/core';

import * as phyllotaxymusic from '../../../phyllotaxymusic/src/public-api'

import { MatSliderModule } from '@angular/material/slider'
import { FormsModule } from '@angular/forms';  // Import FormsModule
import { MatGridListModule } from '@angular/material/grid-list';

import { MatRadioChange, MatRadioModule } from '@angular/material/radio';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatCardModule } from '@angular/material/card';
import { MatCheckboxModule } from '@angular/material/checkbox';

import { AngularSplitModule } from 'angular-split';

import { GongsvgDiagrammingComponent } from '@vendored_components/github.com/fullstack-lang/gongsvg/ng-github.com-fullstack-lang-gongsvg/projects/gongsvgspecific/src/lib/gongsvg-diagramming/gongsvg-diagramming'
import { TreeComponent } from '@vendored_components/github.com/fullstack-lang/gongtree/ng-github.com-fullstack-lang-gongtree/projects/gongtreespecific/src/public-api'
import { GongtoneComponent } from '@vendored_components/github.com/fullstack-lang/gongtone/ng-github.com-fullstack-lang-gongtone/projects/gongtonespecific/src/lib/gongtone/gongtone.component'
import { CursorspecificComponent } from '../../../../../cursor/ng-github.com-thomaspeugeot-phyllotaxymusic-cursor/projects/cursorspecific/src/public-api'

import { CommonModule } from '@angular/common';
import { HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

@Component({
  selector: 'lib-phyllotaxymusicspecific',
  standalone: true,
  imports: [
    CommonModule,
    MatSliderModule,
    MatRadioModule,
    MatCardModule,
    MatCheckboxModule,
    FormsModule,
    MatFormFieldModule,
    AngularSplitModule,
    MatGridListModule,
    GongsvgDiagrammingComponent,
    TreeComponent,
    GongtoneComponent,
    CursorspecificComponent
  ],
  templateUrl: './phyllotaxymusicspecific.component.html',
  styleUrls: ['./phyllotaxymusicspecific.component.css'],
})
export class PhyllotaxymusicspecificComponent implements OnInit {

  private socket: WebSocket | undefined

  readonly checked = model(false);
  readonly indeterminate = model(false);

  inputMatRadio($event: MatRadioChange) {
    let event2: Event = new Event('input');
    this.input(event2)
  }

  input($event: Event) {
    let parameter = this.frontRepo!.array_Parameters[0]

    this.parameterService.updateFront(parameter, this.StacksNames.Phylotaxy).subscribe(
      () => {

      }
    )
  }

  StacksNames = phyllotaxymusic.StacksNames
  StackName = phyllotaxymusic.StacksNames.Phylotaxy
  TreeNames = phyllotaxymusic.TreeNames


  toto: number = 40.0
  rowHeight: string = "50px"

  public frontRepo?: phyllotaxymusic.FrontRepo

  constructor(
    private frontRepoService: phyllotaxymusic.FrontRepoService,

    private parameterService: phyllotaxymusic.ParameterService,
  ) {

  }

  ngOnInit(): void {

    console.log("ngOnInit")

    this.frontRepoService.connectToWebSocket(this.StacksNames.Phylotaxy).subscribe(
      gongtablesFrontRepo => {
        this.frontRepo = gongtablesFrontRepo
      }
    )
  }

  formatLabel(value: number): string {
    if (value >= 1000) {
      return Math.round(value / 1000) + 'k';
    }

    return `${value}`;
  }

  // Check if a specific note is played
  isNotePlayed(encoding: number, rank: number): boolean {
    return (encoding & (1 << rank)) !== 0;
  }

  // Check if a specific note is played
  isNotePlayedWithOffset(encoding: number, rank: number): boolean {
    const nbBeatsInTheme = this.frontRepo!.array_Parameters[0].NbOfBeatsInTheme
    const bruteOffsetIndex = rank - this.frontRepo!.array_Parameters[0].ActualBeatsTemporalShift + nbBeatsInTheme

    const offsetIndex = bruteOffsetIndex % nbBeatsInTheme
    return (encoding & (1 << offsetIndex)) !== 0;
  }



  // Toggle the state of a specific note
  toggleNote(rank: number): void {
    const parameter = this.frontRepo!.array_Parameters[0];
    if (this.isNotePlayed(parameter.ThemeBinaryEncoding, rank)) {
      // Turn off the note
      parameter.ThemeBinaryEncoding &= ~(1 << rank);
      this.parameterService.updateFront(parameter, this.StacksNames.Phylotaxy).subscribe(
        () => {

        }
      )
    } else {
      // Turn on the note
      parameter.ThemeBinaryEncoding |= (1 << rank);
      this.parameterService.updateFront(parameter, this.StacksNames.Phylotaxy).subscribe(
        () => {

        }
      )
    }
  }
}
